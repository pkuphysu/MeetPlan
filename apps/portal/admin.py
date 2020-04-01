from django.contrib import admin
from django.urls import reverse
from django.utils.html import mark_safe, escape

from . import models
# Register your models here.


class FriendLinkAdmin(admin.ModelAdmin):
    list_display = [
        'id',
        'name',
        'url',
        'link_to_img',
        'description'
    ]

    def link_to_img(self, obj):
        link = reverse("admin:filemanager_myimg_change", args=[obj.image_id])
        return mark_safe(f'<a href="{link}">{escape(obj.image.__str__())}</a>')

    link_to_img.short_description = '图片'
    list_per_page = 20
    list_select_related = True
    search_fields = ['name', 'description']


admin.site.register(models.FriendLink, FriendLinkAdmin)
