<script lang="ts" setup>
import { ref } from "vue";
import ReCol from "@/components/ReCol";
import { formRules } from "./utils/rule";
import { FormItemProps, FormProps } from "./utils/types";
import { useI18n } from "vue-i18n";

defineOptions({
  name: "editUserInfo"
});

const props = withDefaults(defineProps<FormProps>(), {
  formInline: () => ({})
});
const { t } = useI18n();

const ruleFormRef = ref();
const newFormInline = ref(props.formInline);

const emit = defineEmits<{
  (e: "handleUpdate", v: FormItemProps): void;
}>();

const genderChoices = [
  {
    key: "",
    label: t("user.gender.hide")
  },
  {
    key: "male",
    label: t("user.gender.male")
  },
  {
    key: "female",
    label: t("user.gender.female")
  }
];

const handleUpdate = (row: FormItemProps) => {
  ruleFormRef.value.validate(valid => {
    if (valid) {
      emit("handleUpdate", row);
    }
  });
};
</script>

<template>
  <el-form
    ref="ruleFormRef"
    :model="newFormInline"
    :rules="formRules"
    label-width="82px"
  >
    <el-row :gutter="30">
      <re-col :sm="24" :value="12" :xs="24">
        <el-form-item :label="t('user.name')" prop="name">
          <el-input v-model="newFormInline.name" clearable />
        </el-form-item>
      </re-col>
      <re-col :sm="24" :value="12" :xs="24">
        <el-form-item :label="t('user.phone')" prop="phoneNumber">
          <el-input v-model="newFormInline.phoneNumber" clearable />
        </el-form-item>
      </re-col>
      <re-col :sm="24" :value="12" :xs="24">
        <el-form-item :label="t('user.email')" prop="email">
          <el-input v-model="newFormInline.email" clearable />
        </el-form-item>
      </re-col>
      <re-col :sm="24" :value="12" :xs="24">
        <el-form-item :label="t('user.gender')" prop="gender">
          <el-select v-model="newFormInline.gender" class="w-full">
            <el-option
              v-for="item in genderChoices"
              :key="item.key"
              :label="item.label"
              :value="item.key"
            />
          </el-select>
        </el-form-item>
      </re-col>
      <re-col :sm="24" :value="12" :xs="24">
        <el-form-item :label="t('user.department')" prop="department">
          <el-input v-model="newFormInline.department" clearable />
        </el-form-item>
      </re-col>
      <template v-if="newFormInline.isTeacher">
        <re-col :sm="24" :value="12" :xs="24">
          <el-form-item :label="t('user.office')" prop="office">
            <el-input v-model="newFormInline.office" clearable />
          </el-form-item>
        </re-col>
        <re-col>
          <el-form-item :label="t('user.introduction')" prop="introduction">
            <el-input
              v-model="newFormInline.introduction"
              type="textarea"
              :rows="5"
            />
          </el-form-item>
        </re-col>
      </template>
      <template v-else>
        <re-col :sm="24" :value="12" :xs="24">
          <el-form-item :label="t('user.grade')" prop="grade">
            <el-input v-model="newFormInline.grade" clearable />
          </el-form-item>
        </re-col>
        <re-col :sm="24" :value="12" :xs="24">
          <el-form-item :label="t('user.major')" prop="major">
            <el-input v-model="newFormInline.major" clearable />
          </el-form-item>
        </re-col>
        <re-col :sm="24" :value="12" :xs="24">
          <el-form-item :label="t('user.dorm')" prop="dorm">
            <el-input v-model="newFormInline.dorm" clearable />
          </el-form-item>
        </re-col>
      </template>
    </el-row>
    <el-form-item>
      <el-popconfirm
        :title="t('buttons.hsconfirmdupdate')"
        @confirm="handleUpdate(newFormInline)"
      >
        <template #reference>
          <el-button>{{ t("buttons.hssave") }}</el-button>
        </template>
      </el-popconfirm>
    </el-form-item>
  </el-form>
</template>
