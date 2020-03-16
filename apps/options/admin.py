from django.contrib import admin
from . import models
# Register your models here.


class OptionAdmin(admin.ModelAdmin):
    list_display = [
        'id',
        'app',
        'name',
        'value',
    ]
    list_filter = ['app']
    search_fields = ['name']
    list_per_page = 20
    list_select_related = True


admin.site.register(models.Option, OptionAdmin)
