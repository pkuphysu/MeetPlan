import { h, onMounted, reactive, type Ref, ref } from "vue";
import { delay, delObjectProperty, deviceDetection } from "@pureadmin/utils";
import editForm from "./form.vue";
import type {
  AdaptiveConfig,
  LoadingConfig,
  PaginationProps
} from "@pureadmin/table";
import {
  createMajor,
  getMajorList,
  type Major,
  updateMajor
} from "@/api/major";
import { addDialog } from "@/components/ReDialog/index";

export function useColumns() {
  // editMap 是用来存储原始数据的
  const editMap: Ref<Record<number, { editable: boolean } & Major>> = ref({});
  const dataList: Ref<Major[]> = ref([]);
  const loading = ref(true);
  const searchForm = reactive({
    major: ""
  });
  const formRef = ref();

  /** 分页配置 */
  const pagination = reactive<PaginationProps>({
    pageSize: 10,
    currentPage: 1,
    pageSizes: [10, 50, 100],
    total: 0,
    align: "right",
    background: true,
    small: false
  });

  /** 加载动画配置 */
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

  /** 撑满内容区自适应高度相关配置 */
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
    editMap.value = {};
    loadingConfig.text = `正在加载...`;
    loading.value = true;
    delay(600).then(() => {
      getTableData();
      loading.value = false;
    });
  }

  const resetForm = formEl => {
    if (!formEl) return;
    formEl.resetFields();
    onSearch();
  };

  function getTableData() {
    getMajorList(pagination.currentPage, pagination.pageSize, searchForm.major)
      .then(res => {
        if (res.code === 0) {
          dataList.value = res.data;
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
    editMap.value = {};
    pagination.currentPage = 1;
    loadingConfig.text = `正在加载第${pagination.currentPage}页...`;
    loading.value = true;
    delay(600).then(() => {
      getTableData();
      loading.value = false;
    });
  }

  function onCurrentChange(val: number) {
    editMap.value = {};
    loadingConfig.text = `正在加载第${val}页...`;
    loading.value = true;
    delay(600).then(() => {
      getTableData();
      loading.value = false;
    });
  }

  const columns: TableColumnList = [
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
      label: "Major",
      prop: "major",
      cellRenderer: ({ row, index }) => (
        <>
          {editMap.value[index]?.editable ? (
            <el-input v-model={row.major} />
          ) : (
            <p>{row.major}</p>
          )}
        </>
      )
    },
    {
      label: "操作",
      fixed: "right",
      slot: "operation"
    }
  ];

  function openDialog() {
    addDialog({
      title: "新建专业",
      props: {
        formInline: {}
      },
      width: "40%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(editForm, { ref: formRef }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = options.props.formInline as Major;

        function chores() {
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }

        FormRef.validate(valid => {
          if (valid) {
            console.log("curData", curData);
            // 表单规则校验通过
            createMajor(curData)
              .then(res => {
                if (res.code === 0) {
                  chores();
                } else {
                  console.log(res.error);
                }
              })
              .catch(err => {
                console.log(err);
              });
          }
        });
      }
    });
  }

  onMounted(() => {
    delay(600).then(() => {
      getTableData();
      loading.value = false;
    });
  });

  function onEdit(row: Major, index: number) {
    console.log("onEdit", row, index);
    editMap.value[index] = Object.assign({ ...row, editable: true });
  }

  function onSave(index: number) {
    updateMajor(dataList.value[index])
      .then(res => {
        if (res.code === 0) {
          editMap.value[index].editable = false;
        } else {
          console.log(res.error);
        }
      })
      .catch(err => {
        console.log(err);
      });
  }

  function onCancel(index: number) {
    editMap.value[index].editable = false;
    dataList.value[index] = delObjectProperty(editMap.value[index], "editable");
  }

  return {
    searchForm,
    onSearch,
    resetForm,
    openDialog,
    editMap,
    columns,
    dataList,
    onEdit,
    onSave,
    onCancel,
    loading,
    loadingConfig,
    pagination,
    adaptiveConfig,
    onSizeChange,
    onCurrentChange
  };
}
