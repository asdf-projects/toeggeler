interface IError {
    error: string;
}

export const getErrorMessage = (responseBodyJson: IError): string => {
    const errorCode = responseBodyJson.error;
    return `Error.${errorCode}`;
};
