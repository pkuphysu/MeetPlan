import { onMounted, reactive, type Ref, ref } from "vue";
import { delay, delObjectProperty } from "@pureadmin/utils";

import type {
  AdaptiveConfig,
  LoadingConfig,
  PaginationProps
} from "@pureadmin/table";
import { getGradeList, type Grade, updateGrade } from "@/api/grade";

export function useColumns() {
  // editMap 是用来存储原始数据的
  const editMap: Ref<Record<number, { editable: boolean } & Grade>> = ref({});
  const dataList: Ref<Grade[]> = ref([]);
  const loading = ref(true);

  /** 分页配置 */
  const pagination = reactive<PaginationProps>({
    pageSize: 10,
    currentPage: 1,
    pageSizes: [10, 20, 40],
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

  function getTableData() {
    getGradeList(pagination.currentPage, pagination.pageSize)
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
      cellRenderer: ({ row }) => (
        <>
          <p>{row.id}</p>
        </>
      )
    },
    {
      label: "Grade",
      prop: "grade",
      cellRenderer: ({ row, index }) => (
        <>
          {editMap.value[index]?.editable ? (
            <el-input v-model={row.grade} />
          ) : (
            <p>{row.grade}</p>
          )}
        </>
      )
    },
    {
      label: "已毕业离校",
      prop: "isGraduated",
      cellRenderer: ({ row, index }) => (
        <>
          {editMap.value[index]?.editable ? (
            <el-switch
              v-model={row.isGraduated}
              inline-prompt
              active-value={true}
              inactive-value={false}
              active-text="是"
              inactive-text="否"
            />
          ) : (
            <p>{row.isGraduated === true ? "是" : "否"}</p>
          )}
        </>
      )
    },
    // {
    //   label: "爱好",
    //   prop: "hobby",
    //   cellRenderer: ({ row, index }) => (
    //     <>
    //       {editMap.value[index]?.editable ? (
    //         <el-select v-model={row.hobby} clearable placeholder="请选择爱好">
    //           {options.map(item => {
    //             return (
    //               <el-option
    //                 key={item.value}
    //                 label={item.label}
    //                 value={item.value}
    //               />
    //             );
    //           })}
    //         </el-select>
    //       ) : (
    //         <el-tag type="primary">
    //           {options.filter(opt => opt.value == row.hobby)[0]?.label}
    //         </el-tag>
    //       )}
    //     </>
    //   )
    // },
    // {
    //   label: "日期",
    //   prop: "date",
    //   cellRenderer: ({ row, index }) => (
    //     <>
    //       {editMap.value[index]?.editable ? (
    //         <el-date-picker
    //           v-model={row.date}
    //           type="date"
    //           format="YYYY/MM/DD"
    //           value-format="YYYY-MM-DD"
    //           placeholder="请选择日期"
    //         />
    //       ) : (
    //         <p>{row.date}</p>
    //       )}
    //     </>
    //   ),
    //   minWidth: 110
    // },
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

  function onEdit(row: Grade, index: number) {
    console.log("onEdit", row, index);
    editMap.value[index] = Object.assign({ ...row, editable: true });
  }

  function onSave(index: number) {
    updateGrade(dataList.value[index])
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
