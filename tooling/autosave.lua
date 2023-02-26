local bufnr = 12
--:echo nvim_get_current_buf()

vim.api.nvim_create_autocmd("BufWritePost", {
  group = vim.api.nvim_create_augroup("SaveGroup", { clear = true}),
  pattern = "thermostat.py",
  callback  = function()
    --vim.fn.jobstart({"ls"}, {
    --vim.fn.jobstart({"rsync", "-rltgD", "--recursive", "~/Code/thermostat", "thermo:~/"}, {
    vim.fn.jobstart({"rsync", "-rltgD", "--recursive", ".", "thermo:~/"}, {
      stdout_buffered = true,
      on_stdout = function(_, data)
        if data then
          vim.api.nvim_buf_set_lines(bufnr, -1, -1, false, data)
        end
      end,
      on_stderr = function(_, data)
        if data then
          vim.api.nvim_buf_set_lines(bufnr, -1, -1, false, data)
        end
      end,

    })
  end,
})
