import "./reset.css";
import { message } from "@/utils/message";
import type { FormItemProps } from "./types";
import { updateUserInfoApi, uploadUserInfoAvatarApi } from "@/api/user";
import { onMounted, reactive, ref } from "vue";
import { useUserStoreHook } from "@/store/modules/user";
import { useI18n } from "vue-i18n";
import { delay } from "@pureadmin/utils";

export function useUserInfo() {
  const { t } = useI18n();
  const loading = ref(true);

  // 上传头像信息
  const currentUserInfo = reactive<FormItemProps>({});

  function handleUpdate(row: FormItemProps) {
    updateUserInfoApi(row)
      .then(res => {
        if (res.code === 0) {
          message(t("results.success"), { type: "success" });
          onSearch();
        } else {
          message(`${t("results.failed")}，${res.error}`, { type: "error" });
        }
      })
      .catch(err => {
        message(`${t("results.failed")}，${err}`, { type: "error" });
      });
  }

  async function onSearch() {
    loading.value = true;
    useUserStoreHook()
      .getSelfInfo()
      .then(res => {
        Object.keys(res).forEach(param => {
          currentUserInfo[param] = res[param];
        });
      });
    delay(500).then(() => {
      loading.value = false;
    });
  }

  /** 上传头像 */
  function handleUpload(info) {
    const avatarFile = new File([info], `${currentUserInfo.id}.png`, {
      type: info.type,
      lastModified: Date.now()
    });
    const data = new FormData();
    data.append("file", avatarFile);
    uploadUserInfoAvatarApi(data).then(res => {
      if (res.code === 0) {
        message(t("results.success"), { type: "success" });
        updateUserInfoApi({
          id: currentUserInfo.id,
          avatar: res.data
        }).then(res => {
          if (res.code === 0) {
            message(t("results.success"), { type: "success" });
            onSearch();
          }
        });
      } else {
        // message(`${t("results.failed")}，${res.detail}`, { type: "error" });
      }
    });
  }

  onMounted(() => {
    onSearch();
  });

  return {
    t,
    currentUserInfo,
    handleUpload,
    handleUpdate
  };
}
