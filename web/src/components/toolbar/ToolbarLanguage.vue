<template>
  <v-menu scroll-y>
    <template v-slot:activator="{ props }">
      <v-btn icon v-bind="props">
        <v-icon color="primary">mdi-translate</v-icon>
      </v-btn>
    </template>
    <v-list elevation="1" nav>
      <v-list-item
        v-for="locale in availableLocales"
        :key="locale.code"
        @click="setLocale(locale.code)"
        density="compact"
        :active="locale.code === current"
      >
        <template v-slot:prepend>
          <Icon :icon="`twemoji:flag-${locale.flag}`" class="mr-2"/>
        </template>

        <v-list-item-title> {{ locale.label }}</v-list-item-title>
      </v-list-item>
    </v-list>
  </v-menu>
</template>
<script setup lang="ts">
import {Icon} from "@iconify/vue";
import {useLocale} from "vuetify";
import {onMounted} from "vue";
import {availableLocales} from "@/plugins/i18n";
import {useLanguageStore} from "@/store/language";

const {current} = useLocale();
console.log(current.value);
const languageStore = useLanguageStore();

onMounted(() => {
  setLocale(languageStore.language)
})

const setLocale = (locale: string) => {
  current.value = locale;
  languageStore.setLanguage(locale);
};
</script>
