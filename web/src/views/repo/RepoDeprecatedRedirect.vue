<template>
  <div />
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';

import useApiClient from '~/compositions/useApiClient';

const props = defineProps<{
  repoOwner: string;
  repoName: string;
}>();
const apiClient = useApiClient();
const route = useRoute();
const router = useRouter();

onMounted(async () => {
  const repo = await apiClient.lookupRepo(props.repoOwner, props.repoName);

  // {
  //   path: ':pipelineId',
  //   redirect: (route) => ({ name: 'repo-pipeline', params: route.params }),
  // },
  // {
  //   path: 'build/:pipelineId',
  //   redirect: (route) => ({ name: 'repo-pipeline', params: route.params }),
  //   children: [
  //     {
  //       path: ':procId?',
  //       redirect: (route) => ({ name: 'repo-pipeline', params: route.params }),
  //     },
  //     {
  //       path: 'changed-files',
  //       redirect: (route) => ({ name: 'repo-pipeline-changed-files', params: route.params }),
  //     },
  //     {
  //       path: 'config',
  //       redirect: (route) => ({ name: 'repo-pipeline-config', params: route.params }),
  //     },
  //   ],
  // },

  // TODO: support pipeline and build routes

  const path = route.path
    .replace(`/repos/${props.repoOwner}/${props.repoName}`, `/repos/${repo?.id}`)
    .replace(`/${props.repoOwner}/${props.repoName}`, `/repos/${repo?.id}`);

  await router.replace({ path });
});
</script>
