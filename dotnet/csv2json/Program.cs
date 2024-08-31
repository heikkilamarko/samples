using CommandLine;
using csv2json;

Parser.Default.ParseArguments<CliOptions>(args).WithParsed(o =>
{
    if (string.IsNullOrWhiteSpace(o.Output))
    {
        o.Output = $"{o.Input}.json";
    }

    new Processor().Process(o.Input, o.Output);
});

public class CliOptions
{
    [Option('i', "input", Required = true, HelpText = "input file")]
    public string Input { get; set; } = string.Empty;

    [Option('o', "output", Required = false, HelpText = "output file")]
    public string? Output { get; set; }
}
