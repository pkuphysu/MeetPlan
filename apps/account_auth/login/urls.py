from django.urls import path

from .views import *

urlpatterns = [
    path('phy/', PHYLoginView.as_view(), name='phy-login'),
    path('phy/auth/', PHYAuthView.as_view(), name='phy-auth'),
]
