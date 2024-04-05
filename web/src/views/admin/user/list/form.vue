<script setup lang="ts">
import { ref } from "vue";
import type { UserInfo } from "@/api/user";

interface FormProps {
  formInline: UserInfo;
}

const props = withDefaults(defineProps<FormProps>(), {
  formInline: () => ({})
});

const ruleFormRef = ref();
const newFormInline = ref(props.formInline);

function getRef() {
  return ruleFormRef.value;
}

defineExpose({ getRef });
</script>

<template>
  <el-form ref="ruleFormRef" :model="newFormInline" label-width="82px">
    <el-form-item label="名称" prop="grade">
      <el-input
        v-model="newFormInline.grade"
        clearable
        placeholder="请输入名称"
      />
    </el-form-item>
    <el-form-item label="已毕业离校" prop="isGraduated">
      <el-switch
        v-model="newFormInline.isGraduated"
        active-text="是"
        inactive-text="否"
        :active-value="true"
        :inactive-value="false"
      />
    </el-form-item>
  </el-form>
</template>
