<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import type { Ref } from 'vue';

interface Post {
  id: number;
  title: string;
  content: string;
  imageUrl: string;
}

const route = useRoute();
const post: Ref<Post | null> = ref(null);

const postId = computed(() => {
  const id = Array.isArray(route.params.id) ? route.params.id[0] : route.params.id;
  return parseInt(id, 10);
});

onMounted(async () => {
  try {
    const response = await fetch(`/api/posts/${postId.value}`);
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    post.value = await response.json();
  } catch (error) {
    console.error('Failed to fetch post:', error);
  }
});
</script>

<template>
  <div v-if="post" class="post-detail">
    <h1>{{ post.title }}</h1>
    <img :src="post.imageUrl" :alt="post.title" />
    <p>{{ post.content }}</p>
    <router-link to="/">返回列表</router-link>
  </div>
  <div v-else>
    <p>正在載入文章...</p>
  </div>
</template>

<style scoped>
.post-detail {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}
img {
  max-width: 100%;
  height: auto;
  margin-bottom: 20px;
}
p {
  line-height: 1.6;
}
</style>
