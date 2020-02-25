from rest_framework import serializers

from apps.account_auth.models import User


class UserSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = (
            'identity_id',
            'user_name',
            'email',
            'is_teacher',
            'is_superuser',
            'is_active'
        )
