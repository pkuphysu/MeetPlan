from django.contrib import admin
from . import models


# Register your models here.


class MeetPlanAdmin(admin.ModelAdmin):
    list_display = [
        'id',
        'teacher',
        'place',
        'start_time',
        'end_time',
        'allow_other',
        'message',
        'available_choice',
    ]
    search_fields = ['teacher']
    list_per_page = 20
    list_select_related = True


admin.site.register(models.MeetPlan, MeetPlanAdmin)


class MeetPlanOrderAdmin(admin.ModelAdmin):
    list_display = [
        'id',
        'meet_plan',
        'student',
        'completed',
        'message'
    ]
    search_fields = ['meet_plan', 'student']
    list_per_page = 20
    list_select_related = True


admin.site.register(models.MeetPlanOrder, MeetPlanOrderAdmin)


class SemesterDateRangeAdmin(admin.ModelAdmin):
    list_display = [
        'id',
        'start_date',
        'end_date',
    ]
    list_per_page = 20


admin.site.register(models.SemesterDateRange, SemesterDateRangeAdmin)


class FeedBackAdmin(admin.ModelAdmin):
    list_display = [
        'id',
        'teacher',
        'message',
    ]
    search_fields = ['teacher']
    list_per_page = 20
    list_select_related = True


admin.site.register(models.FeedBack, FeedBackAdmin)
