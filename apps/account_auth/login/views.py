from django.shortcuts import render
from django.http import HttpResponse, HttpResponseRedirect
from django.template.response import TemplateResponse
from django.views.generic import View
import requests, hashlib
from django.conf import settings
from django.contrib.auth import get_user_model
from django.contrib.auth import login
from django.urls import reverse
from django.conf import settings


class IAAALoginView(View):

    def get(self, request):
        if settings.DEBUG:
            return TemplateResponse(request, template='account_auth/login/local_login.html')

        response = TemplateResponse(request, template='account_auth/login/iaaa_login.html')
        if request.META.get('QUERY_STRING'):
            response.set_cookie(key='next', value=request.META.get('QUERY_STRING'), expires=5 * 60)
        return response

    def post(self, request):
        if settings.DEBUG:
            identity_id = request.POST.get('username', None)
            user_model = get_user_model()
            user = user_model.objects.filter(identity_id=identity_id)
            if user.count():
                login(request, user[0])
                return HttpResponseRedirect(reverse('portal:index'))
            else:
                return HttpResponseRedirect(reverse('account_auth:iaaa_login'))


class IAAALoginAuth(View):
    def get(self, request):
        rand = request.GET.get('rand')
        token = request.GET.get('token')

        if request.META.get('HTTP_X_FORWARDED_FOR'):
            remote_ip = request.META.get('HTTP_X_FORWARDED_FOR')
        else:
            remote_ip = request.META.get('REMOTE_ADDR')

        # remote_ip = '10.128.202.184'
        print(request.META.get('HTTP_X_FORWARDED_FOR'))
        print('REMOTE_ADDR-' + request.META.get('REMOTE_ADDR'))
        print('REMOTE_IP-' + remote_ip)

        PARA_STR = "appId=%s&remoteAddr=%s&token=%s" % (settings.APPID, remote_ip, token) + settings.APPKEY
        print('PARA_STR:' + PARA_STR)

        msgAbs = hashlib.md5()
        msgAbs.update(PARA_STR.encode('utf-8'))
        print(msgAbs)
        print('msgAbs-' + msgAbs.hexdigest())

        url = "https://iaaa.pku.edu.cn/iaaa/svc/token/validate.do?remoteAddr=%s&appId=%s&token=%s&msgAbs=%s" % \
              (remote_ip, settings.APPID, token, msgAbs.hexdigest())
        print(url)
        try:
            iaaa_response = requests.get(url=url)
        except requests.exceptions.ConnectionError:
            return HttpResponse('服务器网络错误, 请稍后重新登录!\n如果您可以联系管理员, 感激不尽!')
        print(iaaa_response.json())

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
                login(request=request, user=user[0])
                if request.COOKIES.get('next'):
                    cookie_next = request.COOKIES.get('next')
                    response = HttpResponseRedirect(cookie_next.split('=')[1])
                    response.delete_cookie('next')
                    return response
                else:
                    return HttpResponseRedirect(reverse('portal:index'))
            else:
                return HttpResponse('errMsg:%s%s%s' % (name, identity_type, '本应用仅对物理学院学生与教职工开放，若您符合上述条件，请联系网站管理员。'))
        else:
            errMsg = iaaa_response.json()['errMsg']
            return HttpResponse('errMsg:%s\n%s' % (errMsg, '请联系网站管理员'))
