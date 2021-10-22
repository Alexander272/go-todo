export const dateFormat = date => {
    const dateOptions = {
        hour: 'numeric',
        minute: 'numeric',
        second: 'numeric',
        weekday: 'long',
        year: '2-digit',
        month: 'short',
        day: 'numeric',
    }
    return new Date(date).toLocaleDateString('RU-ru', dateOptions)
}
