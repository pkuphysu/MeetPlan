<script lang="ts" setup>
import { reactive, type Ref, ref, watch } from "vue";
import ReCol from "@/components/ReCol";
import { formRules } from "./utils/rule";
import { FormItemProps, FormProps } from "./utils/types";
import { useI18n } from "vue-i18n";
import { getGradeList, Grade } from "@/api/grade";
import { getMajorList, Major } from "@/api/major";
import { Department, getDepartmentList } from "@/api/department";
import { debounce } from "@pureadmin/utils";

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
  pageSize: 10,
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

const handleUpdate = (row: FormItemProps) => {
  ruleFormRef.value.validate(valid => {
    if (valid) {
      emit("handleUpdate", row);
    }
  });
};

// 从 props 中获取的数据，需要在 setup 函数中使用
watch(props, val => {
  gradeList.value.push({
    id: props.formInline.gradeID,
    grade: props.formInline.grade
  });
  departmentList.value.push({
    id: props.formInline.departmentID,
    department: props.formInline.department
  });
  majorList.value.push({
    id: props.formInline.majorID,
    major: props.formInline.major
  });
});
</script>

<template>
  <el-form
    ref="ruleFormRef"
    :model="newFormInline"
    :rules="formRules"
    label-width="auto"
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
          <el-select
            v-model="newFormInline.departmentID"
            filterable
            remote
            reserve-keyword
            placeholder="请选择系所"
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
          <el-form-item :label="t('user.major')" prop="major">
            <el-select
              v-model="newFormInline.majorID"
              filterable
              remote
              reserve-keyword
              placeholder="请选择专业"
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
        </re-col>
        <re-col :sm="24" :value="12" :xs="24">
          <el-form-item :label="t('user.grade')" prop="grade">
            <el-select
              v-model="newFormInline.gradeID"
              filterable
              remote
              reserve-keyword
              placeholder="请选择年级"
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
