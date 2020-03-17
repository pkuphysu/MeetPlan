from django import template
from django.db.models import Sum
from django.utils import timezone

from ..utils import get_term_date

register = template.Library()


@register.filter
def in_this_term(meetplan):
    date_range = get_term_date()
    return meetplan.filter(start_time__gte=date_range[0], end_time__lte=date_range[1])


@register.filter
def in_this_term_after_now(meetplan):
    date_range = get_term_date()
    return meetplan.filter(start_time__gte=timezone.now(), end_time__lte=date_range[1])


@register.filter
def available_num(meetplan):
    return meetplan.aggregate(Sum('available_choice'))['available_choice__sum']
