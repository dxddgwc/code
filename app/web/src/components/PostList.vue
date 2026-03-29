<script setup lang="ts">
import { ref, onMounted } from 'vue';
import type { Ref } from 'vue';

interface Post {
  id: number;
  title: string;
  content: string;
  imageUrl: string;
}

const posts: Ref<Post[]> = ref([]);

onMounted(async () => {
  try {
    const response = await fetch('/api/posts');
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    posts.value = await response.json();
  } catch (error) {
    console.error('Failed to fetch posts:', error);
  }
});
</script>

<template>
  <div class="post-list">
    <h1>部落格文章</h1>
    <ul>
      <li v-for="post in posts" :key="post.id">
        <router-link :to="{ name: 'PostDetail', params: { id: post.id } }">
          <h2>{{ post.title }}</h2>
        </router-link>
      </li>
    </ul>
  </div>
</template>

<style scoped>
.post-list {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  margin-bottom: 20px;
  border-bottom: 1px solid #eee;
  padding-bottom: 20px;
}
a {
  text-decoration: none;
  color: #333;
}
a:hover {
  color: #007bff;
}
</style>
