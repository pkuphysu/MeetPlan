from rest_framework import serializers
from .models import MeetPlan, MeetPlanOrder


class MeetPlanSerializers(serializers.ModelSerializer):
    class Meta:
        model = MeetPlan
        fields = [
            'id',
            'tea_id',
            'place',
            'start_time',
            'end_time',
            'allow_other',
        ]


class MeetPlanOrderSerializers(serializers.ModelSerializer):
    class Meta:
        model = MeetPlanOrder
        fields = [
            'id',
            'meet_plan_id',
            'stu_id',
            'completed',
        ]
