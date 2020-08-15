from django.conf import settings


def sys_setting(request):
    setting_dict = dict()
    setting_dict['version'] = settings.VERSIONS
    setting_dict['debug'] = settings.DEBUG
    setting_dict['domain'] = settings.SITE_URL
    return setting_dict
