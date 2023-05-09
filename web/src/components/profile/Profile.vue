<script setup lang="ts">
import {getAvatarUrl} from "@/utils/utils";
import {computed, reactive} from "vue";
import {getUser, User} from "@/api/user";
import {useUserStore} from "@/store/user";
import {getOption} from "@/api/option";

const userStore = useUserStore();

interface Props {
  userid: number
  asAdmin?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  userid: 0,
  asAdmin: false
})

const userRef = reactive<{
  user: User
}>({
  user: {
    id: props.userid,
    email: "",
    is_active: false,
    is_admin: false,
    is_teacher: false,
    name: "",
    pku_id: ""
  }
})

const adminRef = computed(() => {
  return props.asAdmin && userStore.isAdmin
})

getUser(props.userid).then((res) => {
  userRef.user = res
}).catch((err) => {
  console.log(err)
})

const choicesRef = reactive<{
  departments: string[]
  majors: string[]
  grades: string[]

}>({
  departments: [],
  majors: [],
  grades: []
})

if (sessionStorage.getItem('departments')) {
  choicesRef.departments = JSON.parse(sessionStorage.getItem('departments'))
} else {
  getOption('departments').then((res) => {
    choicesRef.departments = JSON.parse(res)
    sessionStorage.setItem('departments', res)
  })
}

if (sessionStorage.getItem('majors')) {
  choicesRef.majors = JSON.parse(sessionStorage.getItem('majors'))
} else {
  getOption('majors').then((res) => {
    choicesRef.majors = JSON.parse(res)
    sessionStorage.setItem('majors', res)
  })
}

if (sessionStorage.getItem('grades')) {
  choicesRef.grades = JSON.parse(sessionStorage.getItem('grades'))
} else {
  getOption('grades').then((res) => {
    choicesRef.grades = JSON.parse(res)
    sessionStorage.setItem('grades', res)
  })
}


const clickAvatar = () => {
  if (!userRef.user.avatar) {
    window.open("https://cn.gravatar.com/")
  }
}

</script>

<template>
  <v-row>
    <v-col class="v-col-12">
      <v-card density="default" variant="elevated">
        <v-card-text class="d-flex">
          <v-tooltip location="bottom">
            <span v-if="!userRef.user.avatar">
              默认使用Gravatar头像
            </span>
            <template v-slot:activator="{ props }">
              <v-avatar v-bind="props" density="default" class="rounded-sm me-6" variant="flat" size="100"
                        @click="clickAvatar">
                <v-img :src="getAvatarUrl(userRef.user, 1000)">
                </v-img>
              </v-avatar>
            </template>
          </v-tooltip>

          <v-form class="d-flex flex-column justify-center gap-3">
            <div class="d-flex flex-wrap gap-2">
              <v-btn density="default" variant="elevated" disabled>
                <span class="d-none d-sm-block">上传新头像</span>
              </v-btn>
              <input type="file" name="file" accept=".jpg,.jpeg,.png" hidden>
              <v-btn density="default" variant="tonal" class="text-error" disabled>
                <span class="d-none d-sm-block">重置</span>
              </v-btn>
            </div>
            <p class="text-xs mb-0">接受 JPG, JPEG和PNG，最大800KB</p>
          </v-form>
        </v-card-text>
        <v-card-text>
          <v-form class="mt-6">
            <v-row>
              <v-col cols="12" md="6">
                <v-text-field
                  label="姓名"
                  v-model="userRef.user.name"
                  density="comfortable"
                  variant="outlined"
                ></v-text-field>
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  label="PKU ID"
                  v-model="userRef.user.pku_id"
                  density="comfortable"
                  variant="outlined"
                  :readonly="!adminRef"
                ></v-text-field>
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  label="邮箱"
                  v-model="userRef.user.email"
                  density="comfortable"
                  variant="outlined"
                ></v-text-field>
              </v-col>
              <v-col cols="12" md="6">
                <v-select
                  label="性别"
                  v-model="userRef.user.gender"
                  density="comfortable"
                  variant="outlined"
                  :items="[{title: '男', value: 0}, {title: '女', value: 1}]"
                  item-title="title"
                  item-value="value"
                ></v-select>
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  label="电话"
                  v-model="userRef.user.phone"
                  density="comfortable"
                  variant="outlined"
                ></v-text-field>
              </v-col>
              <v-col cols="12" md="6">
                <v-select
                  label="系所/办公室"
                  v-model="userRef.user.department"
                  density="comfortable"
                  variant="outlined"
                  :items="choicesRef.departments"
                ></v-select>
              </v-col>
              <v-col v-if="!userRef.user.is_teacher" cols="12" md="6">
                <v-select
                  label="专业"
                  v-model="userRef.user.major"
                  density="comfortable"
                  variant="outlined"
                  :items="choicesRef.majors"
                ></v-select>
              </v-col>
              <v-col v-if="!userRef.user.is_teacher" cols="12" md="6">
                <v-select
                  label="年级"
                  v-model="userRef.user.grade"
                  density="comfortable"
                  variant="outlined"
                  :items="choicesRef.grades"
                ></v-select>
              </v-col>
              <v-col v-if="!userRef.user.is_teacher" cols="12" md="6">
                <v-text-field
                  label="宿舍"
                  v-model="userRef.user.dorm"
                  density="comfortable"
                  variant="outlined"
                ></v-text-field>
              </v-col>
              <v-col cols="12" md="6">
                <v-text-field
                  label="办公室"
                  v-model="userRef.user.office"
                  density="comfortable"
                  variant="outlined"
                ></v-text-field>
              </v-col>
              <v-col v-if="userRef.user.is_teacher" cols="12">
                <v-textarea
                  label="个人简介"
                  v-model="userRef.user.introduction"
                  density="comfortable"
                  variant="outlined"
                ></v-textarea>
              </v-col>
              <template v-if="adminRef">
                <v-col cols="4" md="3">
                  <v-switch
                    label="教师"
                    v-model="userRef.user.is_teacher"
                    density="comfortable"
                    variant="outlined"
                    :readonly="!adminRef"
                  ></v-switch>
                </v-col>
                <v-col cols="4" md="3">
                  <v-switch
                    label="管理员"
                    v-model="userRef.user.is_admin"
                    density="comfortable"
                    variant="outlined"
                    :readonly="!adminRef"
                  ></v-switch>
                </v-col>
                <v-col cols="4" md="3">
                  <v-switch
                    label="账号状态"
                    v-model="userRef.user.is_active"
                    density="comfortable"
                    variant="outlined"
                    :readonly="!adminRef"
                  ></v-switch>
                </v-col>
              </template>
            </v-row>
          </v-form>
        </v-card-text>
      </v-card>
    </v-col>
  </v-row>
</template>

<style scoped>
.text-xs {
  font-size: .75rem;
  line-height: 1rem
}

.text-error {
  color: rgb(255, 76, 81) !important;
}

.gap-2 {
  gap: .5rem;
}

.gap-3 {
  gap: 0.75rem;
}
</style>
