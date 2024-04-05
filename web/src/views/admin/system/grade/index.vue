<script setup lang="ts">
import { useColumns } from "./columns";
import { ref } from "vue";
import { PureTable } from "@pureadmin/table";
import { PureTableBar } from "@/components/RePureTableBar";
import { deviceDetection } from "@pureadmin/utils";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import AddFill from "@iconify-icons/ri/add-circle-line";
import Refresh from "@iconify-icons/ep/refresh";

defineOptions({
  name: "GradeList"
});
const formRef = ref();
const tableRef = ref();

const {
  searchForm,
  resetForm,
  onSearch,
  openDialog,
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
  <div class="main">
    <el-form
      ref="formRef"
      :inline="true"
      :model="searchForm"
      class="search-form bg-bg_color w-[99/100] pl-8 pt-[12px] overflow-auto"
    >
      <el-form-item label="名称：" prop="grade">
        <el-input
          v-model="searchForm.grade"
          placeholder="请输入名称"
          clearable
          class="!w-[180px]"
          @keyup.enter="onSearch"
        />
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          :icon="useRenderIcon('ri:search-line')"
          :loading="loading"
          @click="onSearch"
        >
          搜索
        </el-button>
        <el-button :icon="useRenderIcon(Refresh)" @click="resetForm(formRef)">
          重置
        </el-button>
      </el-form-item>
    </el-form>

    <div :class="['flex', deviceDetection() ? 'flex-wrap' : '']">
      <PureTableBar
        :columns="columns"
        class="w-full"
        style="transition: width 220ms cubic-bezier(0.4, 0, 0.2, 1)"
        title="年级管理"
        @refresh="onSearch"
      >
        <template #buttons>
          <el-button
            type="primary"
            :icon="useRenderIcon(AddFill)"
            @click="openDialog()"
          >
            新增
          </el-button>
        </template>
        <template v-slot="{ size, dynamicColumns }">
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
            :columns="dynamicColumns"
            :size="size"
            :pagination="pagination"
            :paginationSmall="size === 'small'"
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
        </template>
      </PureTableBar>
    </div>
  </div>
</template>

<style scoped lang="scss">
:deep(.el-dropdown-menu__item i) {
  margin: 0;
}

.main-content {
  margin: 24px 24px 0 !important;
}

.search-form {
  :deep(.el-form-item) {
    margin-bottom: 12px;
  }
}
</style>
