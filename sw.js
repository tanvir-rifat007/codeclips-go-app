const CACHE_NAME = "codeclipsV1";

// Install service worker and cache the resources
self.addEventListener("install", async function (event) {
  console.log("Service Worker installing...");
  const urlsToCache = [
    "/",

    "https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.6.0/styles/default.min.css",
    "https://unpkg.com/pwacompat",
    "https://fonts.googleapis.com",
    "https://fonts.gstatic.com",
    "https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200",
  ];
  const cache = await caches.open(CACHE_NAME);
  cache.addAll(urlsToCache);
});

// Serve cached content when offline
// self.addEventListener("fetch", function (event) {
//   console.log("Fetching:", event.request.url);

//   event.respondWith(
//     caches.match(event.request).then(function (response) {
//       return response || fetch(event.request);
//     })
//   );
// });

// Fetch event to implement stale-while-revalidate strategy
self.addEventListener("fetch", (event) => {
  console.log("fetching cache");
  event.respondWith(
    (async () => {
      const cache = await caches.open(CACHE_NAME);

      // from the cache;

      const cachedResponse = await cache.match(event.request);

      // Fetch the latest resource from the network
      const fetchPromise = fetch(event.request)
        .then((networkResponse) => {
          // Update the cache with the latest version
          cache.put(event.request, networkResponse.clone());
          return networkResponse;
        })
        .catch(() => cachedResponse); // In case of network failure, use cached response

      // return cached immediately and update cache in the background
      return cachedResponse || fetchPromise;
    })()
  );
});

// Update service worker
// self.addEventListener("activate", function (event) {
//   console.log("Service Worker activating...");

//   const cacheWhitelist = [CACHE_NAME];
//   event.waitUntil(
//     caches.keys().then(function (cacheNames) {
//       return Promise.all(
//         cacheNames.map(function (cacheName) {
//           if (cacheWhitelist.indexOf(cacheName) === -1) {
//             return caches.delete(cacheName);
//           }
//         })
//       );
//     })
//   );
// });
