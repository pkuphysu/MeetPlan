<script setup lang="ts">
import { ref } from "vue";
import type { Option } from "@/api/option";

interface FormProps {
  formInline: Option;
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
    <el-form-item label="键名" prop="option">
      <el-input
        v-model="newFormInline.name"
        clearable
        placeholder="请输入键名"
      />
    </el-form-item>
    <el-form-item label="键值" prop="value">
      <el-input
        v-model="newFormInline.value"
        clearable
        placeholder="请输入键值"
      />
    </el-form-item>
  </el-form>
</template>
