<script lang="ts" setup>
import { ref } from "vue";
import { useUserInfo } from "./utils/hook";
import editUserInfo from "./edit.vue";
import editUserAvatar from "./avatar.vue";
import dayjs from "dayjs";

defineOptions({
  name: "UserInfo"
});

const { currentUserInfo, handleUpdate, handleUpload, t } = useUserInfo();
const activeTab = ref("userinfo");
</script>

<template>
  <el-row :gutter="24">
    <el-col :lg="8" :md="8" :sm="24" :xl="8" :xs="24">
      <el-card>
        <template v-slot:header>
          <div class="clearfix">
            <span>{{ t("userinfo.userinfo") }}</span>
          </div>
        </template>
        <div>
          <div class="text-center">
            <el-image
              :preview-src-list="Array.of(currentUserInfo.avatar)"
              :src="currentUserInfo.avatar"
              class="h-[120px]"
              fit="cover"
            />
          </div>

          <el-descriptions :column="1" size="large">
            <el-descriptions-item :label="t('user.name')"
              >{{ currentUserInfo.name }}
            </el-descriptions-item>
            <el-descriptions-item :label="t('user.pkuID')"
              >{{ currentUserInfo.pkuID }}
            </el-descriptions-item>
            <el-descriptions-item :label="t('user.department')"
              >{{ currentUserInfo.department }}
            </el-descriptions-item>
            <el-descriptions-item :label="t('user.email')"
              >{{ currentUserInfo.email }}
            </el-descriptions-item>
            <el-descriptions-item :label="t('user.gender')">
              {{ currentUserInfo.gender }}
            </el-descriptions-item>
            <el-descriptions-item :label="t('user.roles')">
              <el-space>
                <el-tag v-if="currentUserInfo.isTeacher"
                  >{{ t("user.teacher") }}
                </el-tag>
                <el-tag v-else>{{ t("user.student") }}</el-tag>
                <el-tag v-if="currentUserInfo.isAdmin">
                  {{ t("user.admin") }}
                </el-tag>
              </el-space>
            </el-descriptions-item>
            <el-descriptions-item :label="t('user.registrationDate')">
              {{
                dayjs(currentUserInfo.createdAt).format("YYYY-MM-DD HH:mm:ss")
              }}
            </el-descriptions-item>
            <template v-if="currentUserInfo.isTeacher">
              <el-descriptions-item :label="t('user.office')"
                >{{ currentUserInfo.office }}
              </el-descriptions-item>
              <el-descriptions-item :label="t('user.introduction')"
                >{{ currentUserInfo.introduction }}
              </el-descriptions-item>
            </template>
            <template v-else>
              <el-descriptions-item :label="t('user.grade')"
                >{{ currentUserInfo.grade }}
              </el-descriptions-item>
              <el-descriptions-item :label="t('user.major')"
                >{{ currentUserInfo.major }}
              </el-descriptions-item>
              <el-descriptions-item :label="t('user.dorm')"
                >{{ currentUserInfo.dorm }}
              </el-descriptions-item>
            </template>
          </el-descriptions>
        </div>
      </el-card>
    </el-col>
    <el-col :lg="16" :md="16" :sm="24" :xl="16" :xs="24">
      <el-card>
        <template v-slot:header>
          <div class="clearfix">
            <span>{{ t("userinfo.updateUserInfo") }}</span>
          </div>
        </template>
        <el-tabs v-model="activeTab">
          <el-tab-pane :label="t('userinfo.basicInfo')" name="userinfo">
            <edit-user-info
              :form-inline="currentUserInfo"
              @handle-update="handleUpdate"
            />
          </el-tab-pane>
          <el-tab-pane :label="t('userinfo.updateAvatar')" name="avatar">
            <edit-user-avatar
              v-if="activeTab === 'avatar'"
              :avatar="currentUserInfo.avatar"
              @handle-update="handleUpload"
            />
          </el-tab-pane>
        </el-tabs>
      </el-card>
    </el-col>
  </el-row>
</template>

<style lang="scss" scoped>
:deep(.el-button:focus-visible) {
  outline: none;
}

.main-content {
  margin: 20px 20px 0 !important;
}
</style>
