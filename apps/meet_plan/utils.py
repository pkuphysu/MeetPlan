import dateutil.parser

from apps.options.models import Option


def get_term_date():
    start = Option.objects.get(app='meet_plan', name='term_start_date').value
    end = Option.objects.get(app='meet_plan', name='term_end_date').value

    start = dateutil.parser.parse(start)
    end = dateutil.parser.parse(end)
    return [start, end]
