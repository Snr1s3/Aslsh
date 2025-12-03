package commands

func Echo(parts []string) string {
    if len(parts) < 2 {
        return ""
    }
    returnString := ""
    for i := 1; i < len(parts); i++ {
        returnString += parts[i]
        if i < len(parts)-1 {
            returnString += " "
        }
    }
    return returnString
}