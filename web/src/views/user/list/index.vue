<script setup lang="ts">
import { useColumns } from "./columns";
import { reactive, type Ref, ref } from "vue";
import { PureTable } from "@pureadmin/table";
import { PureTableBar } from "@/components/RePureTableBar";
import { debounce, deviceDetection } from "@pureadmin/utils";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import AddFill from "@iconify-icons/ri/add-circle-line";
import Refresh from "@iconify-icons/ep/refresh";
import { getGradeList, Grade } from "@/api/grade";
import { getMajorList, Major } from "@/api/major";
import { Department, getDepartmentList } from "@/api/department";

defineOptions({
  name: "UserList"
});
const searchFormRef = ref();
const tableRef = ref();

const gradeList: Ref<Grade[]> = ref([]);
const gradePage = reactive({
  query: undefined,
  pageSize: 10,
  currentPage: 1,
  total: 0
});
const majorList: Ref<Major[]> = ref([]);
const majorPage = reactive({
  query: undefined,
  pageSize: 10,
  currentPage: 1,
  total: 0
});
const departmentList: Ref<Department[]> = ref([]);
const departmentPage = reactive({
  query: undefined,
  pageSize: 2,
  currentPage: 1,
  total: 0
});

const getData = (type: string) => {
  if (type === "grade") {
    getGradeList(
      gradePage.currentPage,
      gradePage.pageSize,
      gradePage.query
    ).then(res => {
      gradeList.value = res.data;
      gradePage.total = res.pageInfo.total;
    });
  } else if (type === "major") {
    getMajorList(
      majorPage.currentPage,
      majorPage.pageSize,
      majorPage.query
    ).then(res => {
      majorList.value = res.data;
      majorPage.total = res.pageInfo.total;
    });
  } else if (type === "department") {
    getDepartmentList(
      departmentPage.currentPage,
      departmentPage.pageSize,
      departmentPage.query
    ).then(res => {
      departmentList.value = res.data;
      departmentPage.total = res.pageInfo.total;
    });
  }
};

const onSearchFormRemoteMethod = (type: string) => {
  return debounce(async (query: string) => {
    switch (type) {
      case "grade":
        if (query === gradePage.query) {
          return;
        }
        gradePage.query = query;
        break;
      case "major":
        if (query === majorPage.query) {
          return;
        }
        majorPage.query = query;
        break;
      case "department":
        if (query === departmentPage.query) {
          return;
        }
        departmentPage.query = query;
        break;
    }
    getData(type);
  }, 500);
};
const onSearchFormPagination = (type: string) => {
  return (page: number) => {
    switch (type) {
      case "grade":
        gradePage.currentPage = page;
        break;
      case "major":
        majorPage.currentPage = page;
        break;
      case "department":
        departmentPage.currentPage = page;
        break;
    }
    getData(type);
  };
};

const {
  loading,
  loadingConfig,
  adaptiveConfig,
  searchForm,
  resetSearchForm,
  onSearch,
  tableColumns,
  tableDataList,
  pagination,
  onSizeChange,
  onCurrentChange
} = useColumns();
</script>

<template>
  <div class="main">
    <el-form
      ref="searchFormRef"
      :inline="true"
      :model="searchForm"
      class="search-form bg-bg_color w-[99/100] pl-8 pt-[12px] overflow-auto"
      label-width="auto"
    >
      <el-form-item label="名字：" prop="name">
        <el-input
          v-model="searchForm.name"
          placeholder="请输入名字"
          clearable
          class="!w-[180px]"
          @keyup.enter="onSearch"
        />
      </el-form-item>
      <el-form-item label="学工号：" prop="pkuID">
        <el-input
          v-model="searchForm.pkuID"
          placeholder="请输入学工号"
          clearable
          class="!w-[180px]"
          @keyup.enter="onSearch"
        />
      </el-form-item>
      <el-form-item label="系所：" prop="departmentID">
        <el-select
          v-model="searchForm.departmentID"
          multiple
          collapse-tags
          collapse-tags-tooltip
          filterable
          remote
          reserve-keyword
          placeholder="请选择系所"
          clearable
          class="!w-[180px]"
          :remote-method="onSearchFormRemoteMethod('department')"
        >
          <el-option
            v-for="item in departmentList"
            :key="item.id"
            :label="item.department"
            :value="item.id"
          />
          <template #footer>
            <el-pagination
              layout="prev, pager, next"
              small
              hide-on-single-page
              :current-page="departmentPage.currentPage"
              :page-size="departmentPage.pageSize"
              :total="departmentPage.total"
              @update:current-page="
                onSearchFormPagination('department')($event)
              "
            />
          </template>
        </el-select>
      </el-form-item>
      <el-form-item label="专业：" prop="majorID">
        <el-select
          v-model="searchForm.majorID"
          multiple
          collapse-tags
          collapse-tags-tooltip
          filterable
          remote
          reserve-keyword
          placeholder="请选择专业"
          clearable
          class="!w-[180px]"
          :remote-method="onSearchFormRemoteMethod('major')"
        >
          <el-option
            v-for="item in majorList"
            :key="item.id"
            :label="item.major"
            :value="item.id"
          />
          <template #footer>
            <el-pagination
              layout="prev, pager, next"
              small
              hide-on-single-page
              :current-page="majorPage.currentPage"
              :page-size="majorPage.pageSize"
              :total="majorPage.total"
              @update:current-page="onSearchFormPagination('major')($event)"
            />
          </template>
        </el-select>
      </el-form-item>
      <el-form-item label="年级：" prop="gradeID">
        <el-select
          v-model="searchForm.gradeID"
          multiple
          collapse-tags
          collapse-tags-tooltip
          filterable
          remote
          reserve-keyword
          placeholder="请选择年级"
          clearable
          class="!w-[180px]"
          :remote-method="onSearchFormRemoteMethod('grade')"
        >
          <el-option
            v-for="item in gradeList"
            :key="item.id"
            :label="item.grade"
            :value="item.id"
          />
          <template #footer>
            <el-pagination
              layout="prev, pager, next"
              small
              hide-on-single-page
              :current-page="gradePage.currentPage"
              :page-size="gradePage.pageSize"
              :total="gradePage.total"
              @update:current-page="onSearchFormPagination('grade')($event)"
            />
          </template>
        </el-select>
      </el-form-item>

      <el-form-item label="账号状态：" prop="isActive">
        <el-select
          v-model="searchForm.isActive"
          placeholder="请选择状态"
          clearable
          class="!w-[180px]"
        >
          <el-option label="启用" :value="true" />
          <el-option label="禁用" :value="false" />
        </el-select>
      </el-form-item>
      <el-form-item label="身份：" prop="isTeacher">
        <el-select
          v-model="searchForm.isTeacher"
          placeholder="请选择身份"
          clearable
          class="!w-[180px]"
        >
          <el-option key="教师" label="教师" :value="true" />
          <el-option label="学生" :value="false" />
        </el-select>
      </el-form-item>
      <el-form-item label="管理员：" prop="isAdmin">
        <el-select
          v-model="searchForm.isAdmin"
          placeholder="请选择"
          clearable
          class="!w-[180px]"
        >
          <el-option label="是" :value="true" />
          <el-option label="否" :value="false" />
        </el-select>
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
        <el-button
          :icon="useRenderIcon(Refresh)"
          @click="resetSearchForm(searchFormRef)"
        >
          重置
        </el-button>
      </el-form-item>
    </el-form>

    <div :class="['flex', deviceDetection() ? 'flex-wrap' : '']">
      <PureTableBar
        :columns="tableColumns"
        class="w-full"
        style="transition: width 220ms cubic-bezier(0.4, 0, 0.2, 1)"
        title="用户管理"
        @refresh="onSearch"
      >
        <template #buttons>
          <el-button type="primary" :icon="useRenderIcon(AddFill)">
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
            :data="tableDataList"
            :columns="dynamicColumns"
            :size="size"
            :pagination="pagination"
            :paginationSmall="size === 'small'"
            @page-size-change="onSizeChange"
            @page-current-change="onCurrentChange"
          />
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
