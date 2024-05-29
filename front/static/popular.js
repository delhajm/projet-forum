document.addEventListener('DOMContentLoaded', () => {
	const wordCountElement = document.getElementById('word-count');
	const wordCountDec = document.getElementById('word-count-dec');
	const wordCountInc = document.getElementById('word-count-inc');
	let wordCount = 0;

	wordCountDec.addEventListener('click', () => {
			if (wordCount > 0) {
					wordCount--;
					wordCountElement.textContent = wordCount;
			}
	});

	wordCountInc.addEventListener('click', () => {
			wordCount++;
			wordCountElement.textContent = wordCount;
	});
});
