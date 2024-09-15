<template>
  <el-container>
    <el-main>
      <el-card v-if="article" class="article-detail">
        <h1>{{ article.Title }}</h1>
        <p>{{ article.Content }}</p>
        <div>
          <el-button type="primary" @click="likeArticle">点赞</el-button>
          <p>点赞数: {{ likes }}</p>
        </div>
      </el-card>
      <div v-else class="no-data">您必须登录/注册才可以阅读文章</div>
    </el-main>
  </el-container>
</template>

<script setup lang="ts">
import { ref, onMounted, onUpdated, onBeforeUpdate, onBeforeMount } from "vue";
import { useRoute } from "vue-router";
import axios from "../axios";
import type { Article, Like } from "../types/Article";

const article = ref<Article | null>(null);
const route = useRoute();
const likes = ref<number>(0)

const { id } = route.params;

const fetchArticle = async () => {
  try {
    const response = await axios.get<Article>(`/articles/${id}`);
    article.value = response.data;
  } catch (error) {
    console.error("Failed to load article:", error);
  }
};

const likeArticle = async () => {
  try {
    const res = await axios.post<Like>(`articles/${id}/like`)
    likes.value = res.data.likes
    await fetchLike()
  } catch (error) {
    console.log('Error Liking article:', error)
  }
};

const fetchLike = async ()=>{
  try{
    const res = await axios.get<Like>(`articles/${id}/like`)
    likes.value = res.data.likes
  }catch(error){
    console.log('Error fetching likes:', error)
  }
}

onMounted(fetchArticle);
onMounted(fetchLike)
</script>

<style scoped>
.article-detail {
  margin: 20px 0;
}

.no-data {
  text-align: center;
  font-size: 1.2em;
  color: #999;
}
</style>
