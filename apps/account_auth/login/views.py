import hashlib

import requests
from django.conf import settings
from django.contrib.auth import get_user_model
from django.contrib.auth import login
from django.core.exceptions import PermissionDenied
from django.http import HttpResponse, HttpResponseRedirect, Http404
from django.template.response import TemplateResponse
from django.urls import reverse
from django.views.generic import View

from ..tasks import send_account_active_email


class IAAALoginView(View):

    def get(self, request):
        if settings.DEBUG:
            return TemplateResponse(request, template='account_auth/login/local_login.html')
        ctx = {
            'app_id': settings.APPID,
            'redirect_url': settings.APPREDIRECTURL
        }
        return TemplateResponse(request, template='account_auth/login/iaaa_login.html', context=ctx)

    def post(self, request):
        if settings.DEBUG:
            identity_id = request.POST.get('username', None)
            user_model = get_user_model()
            user = user_model.objects.filter(identity_id=identity_id)
            if user.count():
                if not user[0].is_active:
                    send_account_active_email.delay(user[0].identity_id)
                    raise PermissionDenied("""<div class="callout callout-success">
                        <h4>验证成功，但您还没有激活账号!</h4>
                        <p>我们已经向您的PKU邮箱发送了一封激活邮件，请注意查收！</p>
                        </div>""")
                login(request, user[0])
                return HttpResponseRedirect(reverse('portal:index'))
            else:
                return HttpResponseRedirect(reverse('account_auth:iaaa_login'))
        else:
            raise Http404()


class IAAALoginAuth(View):
    def get(self, request):
        rand = request.GET.get('rand')
        token = request.GET.get('token')

        if request.META.get('HTTP_X_FORWARDED_FOR'):
            remote_ip = request.META.get('HTTP_X_FORWARDED_FOR')
        else:
            remote_ip = request.META.get('REMOTE_ADDR')

        PARA_STR = "appId=%s&remoteAddr=%s&token=%s" % (settings.APPID, remote_ip, token) + settings.APPKEY

        msgAbs = hashlib.md5()
        msgAbs.update(PARA_STR.encode('utf-8'))

        url = "https://iaaa.pku.edu.cn/iaaa/svc/token/validate.do?remoteAddr=%s&appId=%s&token=%s&msgAbs=%s" % \
              (remote_ip, settings.APPID, token, msgAbs.hexdigest())
        try:
            iaaa_response = requests.get(url=url)
        except requests.exceptions.ConnectionError:
            return HttpResponse('服务器网络错误, 请稍后重新登录!\n如果您可以联系管理员, 感激不尽!')

        status = iaaa_response.json()['success']
        if status:
            user_info = iaaa_response.json()['userInfo']
            identity_id = user_info['identityId']
            name = user_info['name']
            dept_id = user_info['deptId']
            identity_type = user_info['identityType']

            user_model = get_user_model()
            user = user_model.objects.filter(identity_id=identity_id)
            if user.count():
                if user[0].is_active:
                    login(request=request, user=user[0])
                    return HttpResponseRedirect(reverse('portal:index'))
                else:
                    send_account_active_email.delay(user[0].identity_id)
                    raise PermissionDenied("""<div class="callout callout-success">
                    <h4>验证成功，但您还没有激活账号!</h4>
                    <p>我们已经向您的PKU邮箱发送了一封激活邮件，请注意查收！</p>
                    <p>邮件发送可能有延时，请耐心等待～</p>
                    </div>""")
            else:
                raise PermissionDenied(
                    '本应用仅对物理学院学生与教职工开放，若您{}{}符合上述条件，请发送邮件到phyxgb@pku.edu.cn申请注册。'.format(name, identity_id))
        else:
            err_msg = iaaa_response.json()['errMsg']
            return HttpResponse('errMsg:%s\n%s' % (err_msg, '请联系网站管理员'))
