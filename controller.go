package sunday

type Controller func(req Request, resp Response) (res Response, e error)
