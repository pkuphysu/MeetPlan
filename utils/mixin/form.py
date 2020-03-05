class FormMixin:
    """返回错误信息表单"""
    def get_error(self):
        if hasattr(self, 'errors'):
            errors = self.errors.get_json_data()
            print(errors)
            if errors != {}:
                error_tuple = errors.popitem()
                error_list = error_tuple[1]
                error_dict = error_list[0]
                message = error_dict['message']
                print(message)
                return message
            else:
                return None
        else:
            return None

    def as_div(self):
        return self._html_output(
            normal_row='<div class="form-group"> %(label)s %(field)s %(help_text)s</div>',
            error_row='%s',
            row_ender='</div>',
            help_text_html='<p class="help-block">%s</p>',
            errors_on_separate_row=True,
        )