// script.js
document.addEventListener('DOMContentLoaded', () => {
  const searchIcon = document.getElementById('search-icon');
  const searchForm = document.getElementById('search-form');

  searchIcon.addEventListener('click', () => {
      if (searchForm.style.display === 'block') {
          searchForm.style.display = 'none';
      } else {
          searchForm.style.display = 'block';
          searchForm.querySelector('input[type="text"]').focus();
      }
  });

  // Optionally hide the form if clicking outside
  document.addEventListener('click', (event) => {
      if (!searchForm.contains(event.target) && !searchIcon.contains(event.target)) {
          searchForm.style.display = 'none';
      }
  });
});