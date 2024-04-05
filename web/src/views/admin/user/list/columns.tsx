import { onMounted, reactive, type Ref, ref } from "vue";
import { delay } from "@pureadmin/utils";
import type {
  AdaptiveConfig,
  LoadingConfig,
  PaginationProps
} from "@pureadmin/table";
import {
  type FilterParams,
  type QueryUserParams,
  searchUser,
  type UserInfo
} from "@/api/user";

export function useColumns() {
  const tableDataList: Ref<UserInfo[]> = ref([]);
  const loading = ref(true);
  const searchForm: FilterParams = reactive({});

  const pagination = reactive<PaginationProps>({
    pageSize: 10,
    currentPage: 1,
    pageSizes: [10, 50, 100],
    total: 0,
    align: "right",
    background: true,
    small: false
  });

  const loadingConfig = reactive<LoadingConfig>({
    text: "正在加载第一页...",
    viewBox: "-10, -10, 50, 50",
    spinner: `
        <path class="path" d="
          M 30 15
          L 28 17
          M 25.61 25.61
          A 15 15, 0, 0, 1, 15 30
          A 15 15, 0, 1, 1, 27.99 7.5
          L 15 15
        " style="stroke-width: 4px; fill: rgba(0, 0, 0, 0)"/>
      `
    // svg: "",
    // background: rgba()
  });

  const adaptiveConfig: AdaptiveConfig = {
    /** 表格距离页面底部的偏移量，默认值为 `96` */
    offsetBottom: 110
    /** 是否固定表头，默认值为 `true`（如果不想固定表头，fixHeader设置为false并且表格要设置table-layout="auto"） */
    // fixHeader: true
    /** 页面 `resize` 时的防抖时间，默认值为 `60` ms */
    // timeout: 60
    /** 表头的 `z-index`，默认值为 `100` */
    // zIndex: 100
  };

  async function onSearch() {
    console.log("onSearch");
    loadingConfig.text = `正在加载...`;
    loading.value = true;
    delay(600).then(() => {
      getTableData();
      loading.value = false;
    });
  }

  const resetSearchForm = formEl => {
    if (!formEl) return;
    formEl.resetFields();
    onSearch();
  };

  function getTableData() {
    let param: QueryUserParams = {
      page: pagination.currentPage,
      pageSize: pagination.pageSize
    };
    if (searchForm.name) {
      param.name = searchForm.name;
    }
    if (searchForm.pkuID) {
      param.pkuID = searchForm.pkuID;
    }
    if (searchForm.isActive !== undefined) {
      param.isActive = searchForm.isActive;
    }
    if (searchForm.isTeacher !== undefined) {
      param.isTeacher = searchForm.isTeacher;
    }
    if (searchForm.isAdmin !== undefined) {
      param.isAdmin = searchForm.isAdmin;
    }
    if (searchForm.departmentID) {
      param.departmentID = searchForm.departmentID;
    }
    if (searchForm.majorID) {
      param.majorID = searchForm.majorID;
    }
    if (searchForm.gradeID) {
      param.gradeID = searchForm.gradeID;
    }
    searchUser(param)
      .then(res => {
        if (res.code === 0) {
          tableDataList.value = res.data;
          pagination.total = res.pageInfo.total;
          loading.value = false;
        } else {
          console.log(res.error);
        }
      })
      .catch(err => {
        console.log(err);
      });
  }

  function onSizeChange(val: number) {
    console.log("onSizeChange", val);
    pagination.currentPage = 1;
    loadingConfig.text = `正在加载第${pagination.currentPage}页...`;
    loading.value = true;
    delay(600).then(() => {
      getTableData();
      loading.value = false;
    });
  }

  function onCurrentChange(val: number) {
    loadingConfig.text = `正在加载第${val}页...`;
    loading.value = true;
    delay(600).then(() => {
      getTableData();
      loading.value = false;
    });
  }

  const tableColumns: TableColumnList = [
    {
      label: "ID",
      prop: "id",
      fixed: "left",
      minWidth: "80px",
      cellRenderer: ({ row }) => (
        <>
          <p>{row.id}</p>
        </>
      )
    },
    {
      label: "Name",
      prop: "name",
      cellRenderer: ({ row }) => (
        <>
          <p>{row.name}</p>
        </>
      )
    },
    {
      label: "PKU ID",
      prop: "pkuID",
      cellRenderer: ({ row }) => (
        <>
          <p>{row.pkuID}</p>
        </>
      )
    },
    {
      label: "身份",
      prop: "isTeacher",
      cellRenderer: ({ row }) => (
        <>
          <p>{row.isTeacher === true ? "教师" : "学生"}</p>
        </>
      )
    },
    {
      label: "管理员",
      prop: "isAdmin",
      cellRenderer: ({ row }) => (
        <>
          <p>{row.isAdmin === true ? "是" : "否"}</p>
        </>
      )
    },
    {
      label: "账号状态",
      prop: "isActive",
      cellRenderer: ({ row }) => (
        <>
          <p>{row.isActive === true ? "已激活" : "未激活"}</p>
        </>
      )
    },
    {
      label: "邮箱",
      prop: "email",
      cellRenderer: ({ row }) => (
        <>
          <p>{row.email}</p>
        </>
      )
    },
    {
      label: "电话",
      prop: "phoneNumber",
      cellRenderer: ({ row }) => (
        <>
          <p>{row.phoneNumber}</p>
        </>
      )
    },
    {
      label: "操作",
      fixed: "right",
      slot: "operation"
    }
  ];

  onMounted(() => {
    delay(600).then(() => {
      getTableData();
      loading.value = false;
    });
  });

  return {
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
  };
}
