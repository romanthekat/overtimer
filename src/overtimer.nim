import times, os, strutils

type
  EntryTypes {.pure} = enum
    overtime, spending

  Commands = enum
    start, stop, spend, status

  Entry = object
    startTime: DateTime

  FinishedEntry = object
    startTime: DateTime 
    endTime: DateTime 


proc getCommand(): Commands =
  let params = commandLineParams()
  if params.len() > 1:
    raise newException(ValueError, "only one parameter can be specified")
  elif params.len() == 0: 
    return Commands.status

  return parseEnum[Commands](params[0])


when isMainModule:
  let command = getCommand()
  echo command
