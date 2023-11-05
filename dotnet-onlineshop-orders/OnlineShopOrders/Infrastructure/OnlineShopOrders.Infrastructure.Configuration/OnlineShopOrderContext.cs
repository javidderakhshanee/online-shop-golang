using System.Reflection;
using Microsoft.EntityFrameworkCore;
using OnlineShopOrders.Core.Domain;

namespace OnlineShopOrders.Infrastructure.Configuration;

public sealed class OnlineShopOrdersContext : DbContext
{
    public OnlineShopOrdersContext(DbContextOptions options) : base(options)
    {
    }
    
    public DbSet<Order> Orders{get;set;}

    protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
    {
        base.OnConfiguring(optionsBuilder);
    }
    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.ApplyConfigurationsFromAssembly(
            Assembly.GetExecutingAssembly(),
            t => t.GetInterfaces().Any(i => i.IsGenericType &&
                 i.GetGenericTypeDefinition() == typeof(IEntityTypeConfiguration<>) &&
                 (typeof(Entity).IsAssignableFrom(i.GenericTypeArguments[0]) || 
                  typeof(ValueObject).IsAssignableFrom(i.GenericTypeArguments[0])))
        );
    }


}