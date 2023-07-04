<script setup lang="ts">
import useAppStore from "@/pinia/modules/app";
import {inject} from "vue";

// pinia - app
const app = useAppStore()

const a = (a)=>{
  app.theme = a
}
</script>

<template>
  <button @click="a('dark')">dark</button>
  <button @click="a('light')">light</button>
  <button @click="a('auto')">auto</button>

  <router-view></router-view>
</template>

<style lang="scss" scoped>
</style>
