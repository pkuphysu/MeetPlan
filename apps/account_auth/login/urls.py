from django.urls import path
from .views import *

urlpatterns = [
    path('iaaa/', IAAALoginView.as_view(), name='iaaa_login'),
    path('iaaa/auth/', IAAALoginAuth.as_view(), name='iaaa_auth'),
]
