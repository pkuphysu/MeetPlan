from django.views.generic.edit import CreateView
from django import forms
from apps.filemanager.models import File, Img
from utils.mixin.form import FormMixin


class FileUploadForm(forms.ModelForm, FormMixin):
    class Meta:
        model = File
        fields = ['file']
        labels = {}
        widgets = {}


class FileUploadViewMixin(CreateView):
    model = File
    form_class = FileUploadForm

    def form_valid(self, form):
        form.instance.user = self.request.user
        form.instance.upload_or_download = True
        return super().form_valid(form)


class ImgUploadForm(forms.ModelForm, FormMixin):
    class Meta:
        model = Img
        fields = ['img']
        labels = {}
        widgets = {}


class ImgUploadViewMixin(CreateView):
    model = Img
    form_class = ImgUploadForm

    def form_valid(self, form):
        form.instance.user = self.request.user
        form.instance.upload_or_download = True
        return super().form_valid(form)


