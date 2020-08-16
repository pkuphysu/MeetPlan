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
    search_fields = ['teacher__user_name']
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
    search_fields = ['meet_plan__teacher__user_name', 'student__user_name']
    list_per_page = 20
    list_select_related = True


admin.site.register(models.MeetPlanOrder, MeetPlanOrderAdmin)


class FeedBackAdmin(admin.ModelAdmin):
    list_display = [
        'id',
        'teacher',
        'message',
    ]
    search_fields = ['teacher__user_name']
    list_per_page = 20
    list_select_related = True


admin.site.register(models.FeedBack, FeedBackAdmin)
