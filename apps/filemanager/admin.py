from django.contrib import admin

# Register your models here.
from . import models


class MyFileAdmin(admin.ModelAdmin):
    list_display = [
        'user',
        'file',
        'upload_or_generate',
        'create_time'
    ]
    list_filter = ['upload_or_generate']
    search_fields = ['user', 'upload_or_generate']
    list_per_page = 20
    list_select_related = True


admin.site.register(models.MyFile, MyFileAdmin)


class MyImgAdmin(admin.ModelAdmin):
    list_display = [
        'user',
        'img',
        'upload_or_generate',
        'create_time'
    ]
    list_filter = ['upload_or_generate']
    search_fields = ['user']
    list_per_page = 20
    list_select_related = True


admin.site.register(models.MyImg, MyImgAdmin)
