class FormMixin:

    def as_div(self):
        return self._html_output(
            normal_row='<div class="form-group"> %(label)s%(field)s'
                       '<span class="text-red">%(errors)s</span> %(help_text)s</div>',
            error_row='%s',
            row_ender='</div>',
            help_text_html='<p class="help-block">%s</p>',
            errors_on_separate_row=False,
        )
