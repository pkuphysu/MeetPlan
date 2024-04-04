<script setup lang="ts">
import { ref } from "vue";
import type { Major } from "@/api/major";

interface FormProps {
  formInline: Major;
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
    <el-form-item label="名称" prop="major">
      <el-input
        v-model="newFormInline.major"
        clearable
        placeholder="请输入名称"
      />
    </el-form-item>
  </el-form>
</template>
