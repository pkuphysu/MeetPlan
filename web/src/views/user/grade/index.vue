<script setup lang="ts">
import { useColumns } from "./columns";
import { ref } from "vue";
import { PureTable } from "@pureadmin/table";

const tableRef = ref();

const {
  editMap,
  columns,
  dataList,
  onEdit,
  onSave,
  onCancel,
  pagination,
  loading,
  loadingConfig,
  adaptiveConfig,
  onSizeChange,
  onCurrentChange
} = useColumns();
</script>

<template>
  <div class="flex">
    <PureTable
      ref="tableRef"
      border
      adaptive
      :adaptiveConfig="adaptiveConfig"
      row-key="id"
      align-whole="center"
      showOverflowTooltip
      :loading="loading"
      :loading-config="loadingConfig"
      :header-cell-style="{
        background: 'var(--el-fill-color-light)',
        color: 'var(--el-text-color-primary)'
      }"
      :data="dataList"
      :columns="columns"
      :pagination="pagination"
      @page-size-change="onSizeChange"
      @page-current-change="onCurrentChange"
    >
      <template #operation="{ row, index }">
        <el-button
          v-if="!editMap[index]?.editable"
          class="reset-margin"
          link
          type="primary"
          @click="onEdit(row, index)"
        >
          修改
        </el-button>
        <div v-else>
          <el-button
            class="reset-margin"
            link
            type="primary"
            @click="onSave(index)"
          >
            保存
          </el-button>
          <el-button class="reset-margin" link @click="onCancel(index)">
            取消
          </el-button>
        </div>
      </template>
    </PureTable>
  </div>
</template>
