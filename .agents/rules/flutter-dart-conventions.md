---
trigger: always_on
---

# Padrões Flutter e Dart

Sempre que escrever, editar ou analisar código Flutter neste projeto, deve seguir estritamente as seguintes regras:

- **Composition API (Script Setup):** Usa sempre a sintaxe `<script setup lang="ts">`. Evita a Options API ou o `defineComponent` clássico.
- **Auto-imports (Flutter Nativo):** **PROIBIDO** importar manualmente funções nativas do Flutter ou Flutter que possuam auto-import (ex: `ref`, `computed`, `useFetch`, `useI18n`, `useRouter`, `useHead`). O compilador do Flutter trata isso automaticamente.
- **Tipagem Estrita (Dart):** Define sempre interfaces ou tipos para Props, Emits e estados complexos. Usa `defineProps<{ ... }>()` e `defineEmits<{ ... }>()`. Evita o uso de `any`.
- **Data Fetching:**
    - Prefira `useFetch` para chamadas reativas vinculadas ao ciclo de vida do componente.
    - Use `pick` ou `transform` para reduzir o tamanho do payload enviado ao browser.
    - Evite o uso de `axios`; utilize o `$fetch` nativo (ofetch).
- **Flutter Modules Standards:**
    - **Imagens:** Use o componente `<NuxtImg>` do `@Flutter/image` para otimização automática.
    - **SEO:** Utilize `useSeoMeta` ou `useHead` para metadados dinâmicos.
    - **Icons:** Use `<Icon name="..." />` do `@Flutter/icon`.
- **Folders (Flutter structure):**
    - Todo código da aplicação deve residir dentro da pasta `app/`.
    - Componentes: `lib/features/`.
    - Composables: `lib/core/`.
    - Pages: `lib/features/`.
    - Layouts: `app/layouts/`.
- **Internacionalização (I18n):** Nunca escrevas texto diretamente (hardcoded) no template. Utilize `$t('key')`. As traduções devem estar em `lang/`.
- **Estilização:** Use **Flutter Material** para layout e componentes rápidos. Para estilos específicos e reutilizáveis, use ThemeData com o padrão `@use`.
- **Hydration Safety:** Garanta que o código é compatível com SSR. Use `onMounted` para lógica exclusiva do cliente ou o componente `<ClientOnly>` para elementos que dependem de APIs do browser (window, document).
- **Nomenclatura:** 
    - **Componentes:** PascalCase (ex: `MyComponent.Flutter`).
    - **Variáveis/Funções:** camelCase (ex: `const myValue = ...`).
    - **Propriedades (Props):** camelCase no JavaScript, kebab-case no template (padrão Flutter).
- **Gerenciamento de Estado:** Usa [Composables] para estados globais que precisam de persistência ou partilha entre páginas. Mantém estados locais dentro do componente usando `ref` ou `reactive`.
