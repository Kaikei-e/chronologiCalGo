let drinkNum = document.querySelector('.drinkNum');

selectElement.addEventListener('change', (event) => {
  const result = document.querySelector('.Drinks');
  result.textContent = `The number of drinks are ${event.target.value}`;
});
