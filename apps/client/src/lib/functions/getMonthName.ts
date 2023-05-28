export function getMonthName(month: number) {
	const date = new Date();
	date.setMonth(month - 1);
	return date.toLocaleString('en-US', { month: 'long' });
}
