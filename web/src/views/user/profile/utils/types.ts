import type { UserInfo } from "@/api/user";

type FormItemProps = UserInfo;

interface FormProps {
  formInline: FormItemProps;
}

export type { FormItemProps, FormProps };
