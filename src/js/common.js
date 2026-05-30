(function () {
  const nav = document.querySelector('.page-nav');
  const hamburger = document.querySelector('.hamburger-toggle');
  const links = Array.from(document.querySelectorAll('.page-nav a'));
  const sections = Array.from(document.querySelectorAll('main section[id]'));
  const idToLink = {};
  links.forEach(a => {
    const href = a.getAttribute('href') || '';
    if (href.startsWith('#')) idToLink[href.slice(1)] = a;
  });

  // Hamburger menu toggle
  hamburger.addEventListener('click', (ev) => {
    ev.stopPropagation();
    nav.classList.toggle('open');
  });

  // Close menu when clicking outside
  document.addEventListener('click', (ev) => {
    if (!nav.contains(ev.target)) {
      nav.classList.remove('open');
    }
  });

  // Close menu when a link is clicked
  links.forEach(a => {
    a.addEventListener('click', () => {
      nav.classList.remove('open');
    });
  });

  // determine which section should be marked active
  function updateActive() {
    const threshold = window.innerHeight * 0.42;
    let chosen = null;

    const firstSection = sections[0];
    if (firstSection) {
      const firstRect = firstSection.getBoundingClientRect();
      if (firstRect.top > threshold) {
        chosen = 'top';
      }
    }

    if (!chosen) {
      let minDist = Number.POSITIVE_INFINITY;
      sections.forEach(s => {
        const r = s.getBoundingClientRect();
        if (r.top <= threshold && r.bottom > threshold) {
          chosen = s.id;
          minDist = 0;
        } else {
          const dist = Math.min(Math.abs(r.top - threshold), Math.abs(r.bottom - threshold));
          if (dist < minDist) { minDist = dist; chosen = s.id; }
        }
      });
    }

    if (chosen) {
      links.forEach(l => l.classList.remove('active'));
      if (idToLink[chosen]) idToLink[chosen].classList.add('active');
    }
  }

  let raf = null;
  window.addEventListener('scroll', () => {
    if (raf) cancelAnimationFrame(raf);
    raf = requestAnimationFrame(() => { updateActive(); raf = null; });
  }, { passive: true });
  window.addEventListener('resize', updateActive);

  // clicks / hash behavior
  links.forEach(a => a.addEventListener('click', ev => {
    links.forEach(l => l.classList.remove('active'));
    ev.currentTarget.classList.add('active');
  }));
  window.addEventListener('load', () => {
    const h = location.hash.replace('#', '');
    if (h && idToLink[h]) {
      links.forEach(l => l.classList.remove('active'));
      idToLink[h].classList.add('active');
    } else {
      updateActive();
    }
  });
})();
