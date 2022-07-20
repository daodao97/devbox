<template>
  <div class="tools">
    <div v-for="(p, index) in plugins" :key="p.title + '-' + index" class="box">
      <div class="title">
        <span class="title-text">{{ p.title }}</span>
      </div>
      <div class="items">
        <div v-for="(i, j) in p.children" :key="i.title + '-' + j" class="item"
          @click="() => router.push(i.path as string)">
          <div class="item-text">
            {{ i.title }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

const plugins = [
  {
    title: 'Utils',
    children: [
      {
        title: 'Request',
        path: '/request',
        background: '',
      },
      {
        title: 'GRPC',
      },
      {
        title: 'Regex',
        path: '/iframe?title=Regex&src=http://json.cn/'
      },
    ],
  },
  {
    title: 'Transfer',
    children: [
      {
        title: 'JSON',
        path: '/json',
        background: '',
      },
      {
        title: 'localTest',
        path: '/plugin?name=test&title=Test',
      },
      {
        title: 'Crontab',
        path: '/plugin?name=crontab&title=Crontab',
      },
      {
        title: 'Time',
        path: '/iframe?title=Time&src=https://tools.fun/timestamp.html',
      },
      {
        title: 'Bilibili',
        path: '/iframe?title=Time&src=https://bilibili.com',
      },
      {
        title: 'CURL',
      },
    ],
  },
]

</script>

<style lang="scss" scoped>
.tools {
  .box {
    display: flex;
    flex-direction: column;
    margin-bottom: 20px;

    .title {
      margin: 0 10px;

      .title-text {
        display: inline-block;
        height: 30px;
        line-height: 30px;
        border-bottom: 1px solid black;
      }
    }

    .items {
      display: flex;
      flex-direction: row;
      flex-wrap: wrap;

      .item {
        cursor: pointer;
        display: flex;
        flex-direction: column;
        justify-content: space-around;
        padding: 5px;
        width: 80px;
        text-align: center;
        height: 80px;
        border: 1px solid black;
        border-radius: 5px;
        margin: 10px 0 10px 10px;

        .item-text {
          word-wrap: break-word;
        }
      }
    }
  }

}
</style>
