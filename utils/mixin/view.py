from django.views.generic.edit import CreateView
from django import forms
from apps.filemanager.models import File, Img
from utils.mixin.form import FormMixin


class FileUploadForm(forms.ModelForm, FormMixin):
    class Meta:
        model = File
        fields = ['file']
        labels = {
            'file': '文件'
        }


class FileUploadViewMixin(CreateView):
    model = File
    form_class = FileUploadForm

    def form_valid(self, form):
        form.instance.user = self.request.user
        form.instance.upload_or_generate = True
        return super().form_valid(form)


class ImgUploadForm(forms.ModelForm, FormMixin):
    class Meta:
        model = Img
        fields = ['img']
        labels = {
            'img': '图片'
        }


class ImgUploadViewMixin(CreateView):
    model = Img
    form_class = ImgUploadForm

    def form_valid(self, form):
        form.instance.user = self.request.user
        form.instance.upload_or_generate = True
        return super().form_valid(form)


