export const getErrorMessage = (response: Response): string => {
    if (response.status === 400) {
        return 'Error.InvalidInput';
    } else if (response.status === 401) {
        return 'Error.LoginFailed';
    }
    return 'Error.General';
}