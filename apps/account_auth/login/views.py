import requests
from django.conf import settings
from django.contrib.auth import get_user_model
from django.contrib.auth import login
from django.core.exceptions import PermissionDenied
from django.http import HttpResponseRedirect, Http404
from django.urls import reverse
from django.views.generic import View

from ..tasks import send_account_active_email


class PHYLoginView(View):
    def get(self, request):
        return HttpResponseRedirect("https://auth.phy.pku.edu.cn/oidc/authorize/?response_type=code&scope=openid pku&"
                                    f"client_id={settings.CLIENT_ID}&redirect_uri={settings.REDIRECT_URL}")


class PHYAuthView(View):
    def get(self, request):
        code = request.GET.get('code', '')
        if code == '':
            raise Http404()

        res = requests.post(
            settings.TOKEN_ENDPOINT,
            data={
                "code": code,
                "client_id": settings.CLIENT_ID,
                "client_secret": settings.CLIENT_SECRET,
                "grant_type": "authorization_code",
                "redirect_uri": settings.REDIRECT_URL,
            },
        )
        token = res.json()["access_token"]

        res = requests.get(settings.USERINFO_ENDPOINT, headers={"Authorization": f"Bearer {token}"})
        pku_id = res.json()['pku_id']

        user_model = get_user_model()
        user = user_model.objects.filter(identity_id=pku_id)

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
            raise PermissionDenied('本应用仅对物理学院学生与教职工开放，若您符合上述条件，请发送邮件到phyxgb@pku.edu.cn申请注册。')
