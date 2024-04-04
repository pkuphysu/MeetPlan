<script setup lang="ts">
import { ref } from "vue";
import type { Department } from "@/api/department";

interface FormProps {
  formInline: Department;
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
    <el-form-item label="名称" prop="department">
      <el-input
        v-model="newFormInline.department"
        clearable
        placeholder="请输入名称"
      />
    </el-form-item>
  </el-form>
</template>
