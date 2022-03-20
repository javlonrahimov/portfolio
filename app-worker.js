const cacheName = "app-" + "8cd77efcee66f94356e7288e44e30c3a2c632d22";

self.addEventListener("install", event => {
  console.log("installing app worker 8cd77efcee66f94356e7288e44e30c3a2c632d22");

  event.waitUntil(
    caches.open(cacheName).
      then(cache => {
        return cache.addAll([
          "/portfolio",
          "/portfolio/app.css",
          "/portfolio/app.js",
          "/portfolio/manifest.webmanifest",
          "/portfolio/wasm_exec.js",
          "/portfolio/web/app.wasm",
          "/portfolio/web/css/docs.css",
          "/portfolio/web/css/portfolio.css",
          "/portfolio/web/css/prism.css",
          "/portfolio/web/js/prism.js",
          "https://fonts.googleapis.com/css2?family=Montserrat:wght@400;500&display=swap",
          "https://storage.googleapis.com/murlok-github/icon-192.png",
          "https://storage.googleapis.com/murlok-github/icon-512.png",
          
        ]);
      }).
      then(() => {
        self.skipWaiting();
      })
  );
});

self.addEventListener("activate", event => {
  event.waitUntil(
    caches.keys().then(keyList => {
      return Promise.all(
        keyList.map(key => {
          if (key !== cacheName) {
            return caches.delete(key);
          }
        })
      );
    })
  );
  console.log("app worker 8cd77efcee66f94356e7288e44e30c3a2c632d22 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
