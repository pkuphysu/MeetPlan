from .utils import get_term_date


def sys_setting(request):
    date_range = get_term_date()
    setting_dict = {
        'term_start_date': date_range[0].strftime("%Y-%m-%d"),
        'term_end_date': date_range[1].strftime("%Y-%m-%d"),
    }
    return setting_dict
