const CACHE_NAME = "codeclipsV1";
const urlsToCache = [
  "/",
  "/static/css/main.css",
  "/static/images/maskable128.png",
  "/static/icons/maskable192.png",
  "https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.6.0/styles/default.min.css",
  "https://unpkg.com/pwacompat",
  "https://fonts.googleapis.com",
  "https://fonts.gstatic.com",
  "https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200",
];

// Install service worker and cache the resources
self.addEventListener("install", function (event) {
  event.waitUntil(
    caches.open(CACHE_NAME).then(function (cache) {
      return cache.addAll(urlsToCache);
    })
  );
});

// Serve cached content when offline
self.addEventListener("fetch", function (event) {
  event.respondWith(
    caches.match(event.request).then(function (response) {
      return response || fetch(event.request);
    })
  );
});

// Update service worker
self.addEventListener("activate", function (event) {
  const cacheWhitelist = [CACHE_NAME];
  event.waitUntil(
    caches.keys().then(function (cacheNames) {
      return Promise.all(
        cacheNames.map(function (cacheName) {
          if (cacheWhitelist.indexOf(cacheName) === -1) {
            return caches.delete(cacheName);
          }
        })
      );
    })
  );
});
