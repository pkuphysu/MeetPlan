from .models import SemesterDateRange
from django.utils import timezone


def get_semester_date_range():
    semester = SemesterDateRange.objects.last()
    start_date = timezone.datetime(semester.start_date.year,
                                   semester.start_date.month,
                                   semester.start_date.day,
                                   tzinfo=timezone.utc)
    end_date = timezone.datetime(semester.end_date.year,
                                 semester.end_date.month,
                                 semester.end_date.day,
                                 tzinfo=timezone.utc)
    return [start_date, end_date]
